import PropTypes from 'prop-types'
import React, { Component } from 'react'
import { Scrollbars as S } from 'react-custom-scrollbars'
import { connect } from 'react-redux'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import 'regenerator-runtime/runtime'
import styled, { ThemeProvider } from 'styled-components'
import { Container as C } from '@material-ui/core'

import { dimensions } from '@app/constants'
import { AppBar, Drawer } from '@app/containers'
import { app as appSlice } from '@app/slices'
import { base as baseTheme } from '@app/themes'
import { Home, Sura } from './pages'

const Container = styled(C)`
  height: 100%;
`

const Page = styled.div`
  display: flex;
  flex: 1;
  padding-left: ${(props) => props.paddingLeft}px;
  padding-top: ${dimensions.APP_BAR_HEIGHT}px;
`

const Scrollbars = styled(S)`
  flex: 1;
`

class App extends Component {
  static propTypes = {
    isMobile: PropTypes.bool.isRequired,
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
      <ThemeProvider theme={baseTheme}>
        <BrowserRouter>
          <AppBar />
          <Drawer />
          <Switch>
            <Page paddingLeft={paddingLeft}>
              <Scrollbars>
                <Container maxWidth='sm'>
                  <Route path='/sura/:number'>
                    <Sura />
                  </Route>
                  <Route exact path='/'>
                    <Home />
                  </Route>
                </Container>
              </Scrollbars>
            </Page>
          </Switch>
        </BrowserRouter>
      </ThemeProvider>
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
