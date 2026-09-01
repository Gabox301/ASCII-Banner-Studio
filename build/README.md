# Directorio de build

El directorio `build` se utiliza para alojar todos los archivos y assets de compilación de tu aplicación.

La estructura es:

- `bin` – directorio de salida
- `darwin` – archivos específicos de macOS
- `windows` – archivos específicos de Windows

## Mac

El directorio `darwin` contiene los archivos específicos de las compilaciones para Mac.
Pueden personalizarse y usarse como parte del build. Para volver a dejarlos en su estado por defecto, simplemente
eliminalos y vuelve a compilar con `wails build`.

El directorio contiene los siguientes archivos:

- `Info.plist` – el archivo plist principal utilizado para las compilaciones de Mac. Se usa al compilar con `wails build`.
- `Info.dev.plist` – igual que el plist principal, pero se utiliza al compilar con `wails dev`.

## Windows

El directorio `windows` contiene el manifiesto y los archivos rc que se usan al compilar con `wails build`.
Pueden personalizarse según tu aplicación. Para volver a dejarlos en su estado por defecto, simplemente eliminalos y
vuelve a compilar con `wails build`.

- `icon.ico` – el ícono de la aplicación. Se usa al compilar con `wails build`. Si deseas usar un ícono distinto,
  simplemente reemplaza este archivo por el tuyo. Si falta, se generará un nuevo `icon.ico` a partir del archivo
  `appicon.png` del directorio `build`.
- `installer/*` – los archivos utilizados para crear el instalador de Windows. Se usan al compilar con `wails build`.
- `info.json` – los detalles de la aplicación para las compilaciones de Windows. Estos datos los utilizará el
  instalador de Windows, así como la propia aplicación (clic derecho en el exe → propiedades → detalles).
- `wails.exe.manifest` – el archivo de manifiesto principal de la aplicación.
