// Modules to control application life and create native browser window
const { log } = require('console')
const { app, BrowserWindow, ipcMain } = require('electron')
const path = require('path')

if (require('electron-squirrel-startup')) app.quit();

const isDevEnvironment = process.env.DEV_ENV === 'true'

// enable live reload for electron in dev mode
if (isDevEnvironment) {
    require('electron-reload')(__dirname, {
        electron: path.join(__dirname, '..', 'node_modules', '.bin', 'electron'),
        hardResetMethod: 'exit'
    });
}

let mainWindow;

const createWindow = () => {
    
    // Create the browser window.
    mainWindow = new BrowserWindow({
        width: 460,
        height: 480,
        minWidth: 400,
        maxWidth: 500,
        minHeight: 440,
        maxHeight: 680,
        titleBarStyle: "hidden",
        ...(process.platform !== 'darwin' ? { titleBarOverlay: false } : {}),
        webPreferences: {
            preload: path.join(__dirname, 'preload.js')
        }
    })

    mainWindow.setMenu(null);

    // define how electron will load the app
    if (isDevEnvironment) {

        // if your vite app is running on a different port, change it here
        mainWindow.loadURL('http://localhost:5173/');

        mainWindow.webContents.on('before-input-event', (_, input) => {
            if (input.type === 'keyDown' && input.key === 'F12') {
              mainWindow.webContents.openDevTools({ mode: 'detach' });
            }
        });

        log('Electron running in dev mode: 🧪')

    } else {
        
        // when not in dev mode, load the build file instead
        mainWindow.loadFile(path.join(__dirname, 'build', 'index.html'));

        log('Electron running in prod mode: 🚀')
    }
}

const windowSetSize = BrowserWindow.prototype.setSize;
BrowserWindow.prototype.setSize = function (
	width,
	height,
	animate
) {
	if (animate) {
		const [startWidth, startHeight] = this.getSize();
		const [targetWidth, targetHeight] = [width, height];

		const duration = 200;
		const easing = (t, b, c, d) => (t == d ? b + c : c * (-Math.pow(2, (-10 * t) / d) + 1) + b);

		let currentFrame = 0;
		const updateSize = () => {
			currentFrame++;
			windowSetSize.apply(this, [
				Math.round(easing(currentFrame, startWidth, targetWidth - startWidth, duration)),
				Math.round(easing(currentFrame, startHeight, targetHeight - startHeight, duration)),
			]);
			if (currentFrame < duration) setImmediate(updateSize);
		};
		setImmediate(updateSize);
	} else {
		windowSetSize.apply(this, [width, height]);
	}
};

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', createWindow);

app.on('activate', () => {
    // On macOS it's common to re-create a window in the app when the
    // dock icon is clicked and there are no other windows open.
    if (BrowserWindow.getAllWindows().length === 0) createWindow()
})

ipcMain.on("minimize", () => {
    mainWindow.minimize();
});

ipcMain.on("maximize", () => {
    mainWindow.maximize();
});
ipcMain.on("unmaximize", () => {
    mainWindow.unmaximize();
});

ipcMain.on("isMaximized", () => {
    ipcMain.emit("isMaximizedResponse", mainWindow.isMaximized());
});

ipcMain.on("analyze", () => {
    mainWindow.setResizable(false);
    mainWindow.setSize(1100, 700, true);
    setTimeout(() => {
        mainWindow.setResizable(true);
        mainWindow.setMinimumSize(1100, 700);
        mainWindow.setMaximumSize(1920, 1200);
        mainWindow.setSize(1100, 700);
    }, 1000);
})

ipcMain.on("close", () => {
    mainWindow.close();
});

// Quit when all windows are closed, except on macOS. There, it's common
// for applications and their menu bar to stay active until the user quits
// explicitly with Cmd + Q.
// app.on('window-all-closed', () => {
//     if (process.platform !== 'darwin') app.quit()
// })

// In this file you can include the rest of your app's specific main process
// code. You can also put them in separate files and require them here.