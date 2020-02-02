import PropTypes from 'prop-types'
import React from 'react'
import { Scrollbars as S } from 'react-custom-scrollbars'
import ReactGA from 'react-ga'
import { Helmet } from 'react-helmet'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import 'regenerator-runtime/runtime'
import styled, { ThemeProvider } from 'styled-components'
import { Container as C } from '@material-ui/core'

import { dimensions } from '@app/constants'
import { AppBar, Drawer } from '@app/containers'
import favicon from '@app/images/favicon.png'
import { Home, Sura } from '@app/pages'
import { base as baseTheme } from '@app/themes'

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

export default class Component extends React.Component {
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
          <Helmet defaultTitle='Koran'>
            <link href={favicon} rel="icon" type='image/png' />
          </Helmet>
          <AppBar />
          <Drawer />
          <Switch>
            <Page paddingLeft={paddingLeft}>
              <Scrollbars>
                <Container maxWidth='sm'>
                  <Route
                    path='/sura/:number'
                    render={() => {
                      this.trackPage()
                      return (<Sura />)
                    }}
                  />
                  <Route
                    exact
                    path='/'
                    render={() => {
                      this.trackPage()
                      return (<Home />)
                    }}
                  />
                </Container>
              </Scrollbars>
            </Page>
          </Switch>
        </BrowserRouter>
      </ThemeProvider>
    )
  }

  trackPage = () => {
    ReactGA.pageview(window.location.pathname + window.location.search)
  }
}
