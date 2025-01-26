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
        close: () => ipcRenderer.send('close')
    }
)