import PropTypes from 'prop-types'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import 'regenerator-runtime/runtime'
import styled from 'styled-components'
import { Container } from '@material-ui/core'
import { dimensions } from '@app/constants'
import { AppBar, Drawer } from '@app/containers'
import { app as appSlice } from '@app/slices'
import { Quran, Sura } from './pages'

const Page = styled.div`
  padding-left: ${(props) => props.paddingLeft}px;
  padding-top: ${dimensions.APP_BAR_HEIGHT}px;
`

class App extends Component {
  static propTypes = {
    setWidth: PropTypes.func.isRequired,
    width: PropTypes.number
  }

  render () {
    const clientWidth = document.documentElement.clientWidth
    if (clientWidth !== this.props.width) {
      this.props.setWidth(clientWidth)
    }
    const paddingLeft = this.props.isMobile ? 0 : dimensions.LENGTH_200
    return (
      <BrowserRouter>
        <AppBar />
        <Drawer />
        <Switch>
          <Page paddingLeft={paddingLeft}>
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

export default connect(
  (state) => {
    const isMobile = state.app.isMobile
    const width = state.app.width
    return {
      isMobile,
      width
    }
  },
  (dispatch) => {
    return {
      setWidth: (clientWidth) => dispatch(appSlice.actions.setWidth(clientWidth))
    }
  }
)(App)
