import React from 'react'
import ReactDOM from 'react-dom'
import ReactGA from 'react-ga'
import { Provider } from 'react-redux'

import { App, UserDialog } from '@app/containers'
import configureStore from './configureStore'

ReactGA.initialize(process.env.GA_TRACKING_ID)
ReactGA.pageview(window.location.pathname + window.location.search)

ReactDOM.render(
  <Provider store={configureStore()}>
    <App />
    <UserDialog />
  </Provider>,
  document.getElementById('root')
)
