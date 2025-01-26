const { contextBridge, ipcRenderer } = require('electron')

contextBridge.exposeInMainWorld(
    'electron',
    {
        minimize: () => ipcRenderer.send('minimize'),
        maximize: () => ipcRenderer.send('maximize'),
        unmaximize: () => ipcRenderer.send('unmaximize'),
        startAnalyze: () => ipcRenderer.send('analyze'),
        isMaximized: () => {
            return new Promise((res, rej) => {
                ipcRenderer.send('isMaximized');

                ipcRenderer.once("isMaximizedResponse", res);
            })
        },
        close: () => ipcRenderer.send('close'),
        onProgressChanged: (callback) => {
            console.log("1")
            ipcRenderer.on('progress', (_event, value) => {
                console.log("2")
                callback(value)
            })
        },

        onReceivedBundle: (callback) => ipcRenderer.on("bundle", (_event, value) => {
            console.log("3")
            callback(value)
        })
    }
)