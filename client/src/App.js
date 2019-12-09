import React, { Component } from 'react'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import { Quran, Sura } from './pages'

export default class App extends Component {
  render () {
    return (
      <BrowserRouter>
        <Switch>
          <Route children={<Quran />} exact path='/' />
          <Route children={<Sura />} path='/sura/:number' />
        </Switch>
      </BrowserRouter>
    )
  }
}
