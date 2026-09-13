// ASCII Banner Studio — frontend (vanilla JS, sin build con Node).
//
// Carga los bindings que Wails genera en ./wailsjs/go/main/App.js y arma la
// interacción completa: vista previa en vivo, configuración y exportación.

// Cargamos los bindings de forma dinámica para poder mostrar un mensaje
// amigable si alguien abre la app sin haber generado los bindings antes.
let api = null;
try {
  api = await import('./wailsjs/go/main/App.js');
} catch (err) {
  document.getElementById('status').textContent =
    'No se encontraron los bindings de Wails. Ejecutá `wails dev` o `wails build` una vez y volvé a abrir.';
}

if (api === null) {
  throw new Error('Bindings de Wails no disponibles');
}

const { Generate, Export, ListFonts, ListExportFormats, SaveBanner } = api;

// Extensiones sugeridas según el formato de exportación.
const EXT_BY_FORMAT = {
  txt: 'txt',
  javascript: 'js',
  typescript: 'ts',
  rust: 'rs',
  python: 'py',
  json: 'json',
  go: 'go',
  java: 'java',
  csharp: 'cs',
  c: 'c',
  kotlin: 'kt',
  swift: 'swift',
  ruby: 'rb',
  php: 'php',
  dart: 'dart',
  lua: 'lua',
  shell: 'sh',
  powershell: 'ps1',
};

const els = {
  textInput: document.getElementById('text-input'),
  fontSelect: document.getElementById('font-select'),
  spacing: document.getElementById('spacing'),
  spacingValue: document.getElementById('spacing-value'),
  uppercase: document.getElementById('uppercase'),
  trim: document.getElementById('trim'),
  formatSelect: document.getElementById('format-select'),
  output: document.getElementById('output'),
  meta: document.getElementById('meta'),
  status: document.getElementById('status'),
  copyBtn: document.getElementById('copy-btn'),
  saveBtn: document.getElementById('save-btn'),
};

const state = {
  fonts: [],
  font: 'block',
  spacing: 1,
  uppercase: false,
  trim: false,
  format: 'txt',
};

let debounceTimer = null;
let refreshSeq = 0;
// ===== Inicialización =====

async function init() {
  try {
    const [fonts, formats] = await Promise.all([ListFonts(), ListExportFormats()]);
    state.fonts = fonts;
    renderFontSelect(fonts);
    renderFormatSelect(formats);
  } catch (err) {
    showError(err);
  }
  bindEvents();
  await refresh();
}

function renderFontSelect(fonts) {
  els.fontSelect.innerHTML = '';
  for (const font of fonts) {
    const opt = document.createElement('option');
    opt.value = font.ID;
    opt.textContent = font.Name;
    opt.title = font.ID;
    if (font.ID === state.font) opt.selected = true;
    els.fontSelect.appendChild(opt);
  }
  // Si la fuente por defecto ya no existe (p.ej. cambió el registry),
  // usar la primera disponible.
  if (![...els.fontSelect.options].some((o) => o.value === state.font) && els.fontSelect.options.length > 0) {
    state.font = els.fontSelect.options[0].value;
    els.fontSelect.value = state.font;
  }
}

function renderFormatSelect(formats) {
  els.formatSelect.innerHTML = '';
  for (const fmt of formats) {
    const opt = document.createElement('option');
    opt.value = fmt.ID;
    opt.textContent = fmt.Name;
    els.formatSelect.appendChild(opt);
  }
}

function bindEvents() {
  els.textInput.addEventListener('input', scheduleRefresh);

  els.fontSelect.addEventListener('change', () => {
    state.font = els.fontSelect.value;
    scheduleRefresh();
  });

  els.spacing.addEventListener('input', () => {
    state.spacing = Number(els.spacing.value);
    els.spacingValue.textContent = String(state.spacing);
    scheduleRefresh();
  });

  els.uppercase.addEventListener('change', () => {
    state.uppercase = els.uppercase.checked;
    scheduleRefresh();
  });

  els.trim.addEventListener('change', () => {
    state.trim = els.trim.checked;
    scheduleRefresh();
  });

  els.formatSelect.addEventListener('change', () => {
    state.format = els.formatSelect.value;
    scheduleRefresh();
  });

  els.copyBtn.addEventListener('click', copyToClipboard);
  els.saveBtn.addEventListener('click', saveToFile);

  document.addEventListener('keydown', (e) => {
    if ((e.ctrlKey || e.metaKey) && e.key === 'Enter') {
      e.preventDefault();
      refresh();
    }
    if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === 'c' && e.target === els.output) {
      copyToClipboard();
    }
  });
}
// ===== Generación en vivo =====

function options() {
  return {
    Font: state.font,
    Spacing: state.spacing,
    Align: 'left', // la UI no expone alineación; se usa el default del motor
    Uppercase: state.uppercase,
    Trim: state.trim,
  };
}

function scheduleRefresh() {
  clearTimeout(debounceTimer);
  debounceTimer = setTimeout(refresh, 120);
}

async function refresh() {
  const seq = ++refreshSeq;
  const text = els.textInput.value;

  try {
    const isPlain = state.format === '' || state.format === 'txt';
    const output = isPlain ? await Generate(text, options()) : await Export(text, options(), state.format);

    if (seq !== refreshSeq) return; // una petición más nueva ya ganó

    els.output.textContent = output;
    const fmtLabel = els.formatSelect.selectedOptions[0]?.textContent ?? state.format;
    els.meta.textContent = `${fmtLabel} · ${measure(output)}`;
    els.status.textContent = '';
    els.status.classList.remove('ok');
  } catch (err) {
    if (seq !== refreshSeq) return;
    els.output.textContent = '';
    els.meta.textContent = '—';
    showError(typeof err === 'string' ? err : String((err && err.message) || err));
  }
}

// Cuenta líneas y ancho máximo para la metadata de la vista previa.
function measure(output) {
  const lines = output === '' ? [] : output.split('\n');
  let max = 0;
  for (const line of lines) {
    if (line.length > max) max = line.length;
  }
  return `${lines.length} líneas · ${max} caracteres`;
}

function showError(message) {
  els.status.textContent = message;
  els.status.classList.remove('ok');
}

function flashOk(message, ms = 2500) {
  els.status.textContent = message;
  els.status.classList.add('ok');
  clearTimeout(flashOk._t);
  flashOk._t = setTimeout(() => {
    els.status.classList.remove('ok');
    if (els.status.textContent === message) els.status.textContent = '';
  }, ms);
}

// ===== Acciones =====

async function copyToClipboard() {
  const text = els.output.textContent;
  if (!text) return;

  try {
    await navigator.clipboard.writeText(text);
    flashOk('Copiado al portapapeles ✓');
  } catch (err) {
    // Fallback: selección en un textarea oculto.
    const ta = document.createElement('textarea');
    ta.value = text;
    ta.style.position = 'fixed';
    ta.style.opacity = '0';
    document.body.appendChild(ta);
    ta.select();
    try {
      document.execCommand('copy');
      flashOk('Copiado al portapapeles ✓');
    } catch (copyErr) {
      showError('No se pudo copiar: ' + String(copyErr));
    } finally {
      document.body.removeChild(ta);
    }
  }
}

async function saveToFile() {
  const content = els.output.textContent;
  if (!content) {
    showError('No hay contenido para guardar. Escribí un texto primero.');
    return;
  }
  const ext = EXT_BY_FORMAT[state.format] ?? 'txt';
  const stamp = new Date().toISOString().slice(0, 19).replace(/[:T]/g, '-');
  const suggested = `banner-${stamp}.${ext}`;
  try {
    const path = await SaveBanner(content, suggested);
    if (path) flashOk(`Guardado en: ${path}`);
    // Si path es vacío, el usuario canceló el diálogo: no mostramos nada.
  } catch (err) {
    showError(typeof err === 'string' ? err : String((err && err.message) || err));
  }
}

// ===== Arranque =====

init().catch((err) => showError(String(err)));
