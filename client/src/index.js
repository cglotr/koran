import React from 'react'
import ReactDOM from 'react-dom'
import { Provider } from 'react-redux'

import { App, UserDialog } from '@app/containers'
import configureStore from './configureStore'

ReactDOM.render(
  <Provider store={configureStore()}>
    <App />
    <UserDialog />
  </Provider>,
  document.getElementById('root')
)
