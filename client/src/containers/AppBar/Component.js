import PropTypes from 'prop-types'
import React from 'react'
import styled from 'styled-components'
import {
  AppBar as AB,
  IconButton,
  Toolbar,
  Typography
} from '@material-ui/core'
import {
  Menu as MenuIcon,
  Person as PersonIcon,
  PersonOutline as PersonOutlineIcon
} from '@material-ui/icons'

import { LinkRoute, Row } from '@app/components'
import { dimensions } from '@app/constants'

const AppBar = styled(AB)`
  &.MuiAppBar-root {
    background-color: ${(props) => props.theme.colors.primary}
    color: ${(props) => props.theme.colors.font};
    display: flex;
    height: ${dimensions.APP_BAR_HEIGHT}px;
    justify-content: center;
    z-index: 9999;
  }
`

export default class Component extends React.Component {
  static propTypes = {
    isDrawerOpen: PropTypes.bool.isRequired,
    requestSignIn: PropTypes.func.isRequired,
    setIsDrawerOpen: PropTypes.func.isRequired,
    userEmail: PropTypes.string
  }

  render () {
    return (
      <AppBar position='fixed'>
        <Toolbar>
          <Row>
            <IconButton
              color='inherit'
              edge='start'
              onClick={this.handleClick}
            >
              <MenuIcon />
            </IconButton>
            <LinkRoute onClick={this.handleToRootClick} to='/'>
              <Typography noWrap variant='h6'>Koran</Typography>
            </LinkRoute>
          </Row>
          <Row justifyContent='flex-end'>
            <IconButton
              color='inherit'
              onClick={this.handleSignInClick}
            >
              {this.props.userEmail ? <PersonIcon /> : <PersonOutlineIcon />}
            </IconButton>
          </Row>
        </Toolbar>
      </AppBar>
    )
  }

  handleClick = () => {
    this.props.setIsDrawerOpen(!this.props.isDrawerOpen)
  }

  handleSignInClick = () => {
    if (this.props.userEmail) return
    this.props.requestSignIn()
  }

  handleToRootClick = () => {
    this.props.setIsDrawerOpen(false)
  }
}
