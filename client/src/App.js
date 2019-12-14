import React, { Component } from 'react'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import 'regenerator-runtime/runtime'
import { Drawer } from '@app/components'
import { Quran, Sura } from './pages'

export default class App extends Component {
  render () {
    return (
      <>
        <BrowserRouter>
          <Drawer />
          <Switch>
            <Route exact path='/'>
              <Quran />
            </Route>
            <Route path='/sura/:number'>
              <Sura />
            </Route>
          </Switch>
        </BrowserRouter>
      </>
    )
  }
}
