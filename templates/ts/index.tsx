import React from 'react'
import {createRoot} from 'react-dom/client'

const App = () => {
    return (
        <header>
            <h1>Hello World</h1>
        </header>
    )
}

const container = document.querySelector('#app')!
const root = createRoot(container)
root.render(<App/>)