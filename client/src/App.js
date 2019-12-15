import React, { Component } from 'react'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import 'regenerator-runtime/runtime'
import styled from 'styled-components'
import Container from '@material-ui/core/Container'
import { Drawer } from '@app/components'
import { dimensions } from '@app/constants'
import { Quran, Sura } from './pages'

const Page = styled.div`
  padding-left: ${dimensions.LENGTH_200}px;
`

export default class App extends Component {
  render () {
    return (
      <BrowserRouter>
        <Drawer />
        <Switch>
          <Page>
            <Container maxWidth='md'>
              <Route exact path='/'>
                <Quran />
              </Route>
              <Route path='/sura/:number'>
                <Sura />
              </Route>
            </Container>
          </Page>
        </Switch>
      </BrowserRouter>
    )
  }
}
