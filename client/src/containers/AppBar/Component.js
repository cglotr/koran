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
    isSignedIn: PropTypes.bool.isRequired,
    requestSignIn: PropTypes.func.isRequired,
    setIsDrawerOpen: PropTypes.func.isRequired,
    setIsUserDialogOpen: PropTypes.func.isRequired,
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
              onClick={this.handleProfileClick}
            >
              {this.renderProfileIcon()}
            </IconButton>
          </Row>
        </Toolbar>
      </AppBar>
    )
  }

  handleClick = () => {
    this.props.setIsDrawerOpen(!this.props.isDrawerOpen)
  }

  handleProfileClick = () => {
    if (!this.props.isSignedIn) {
      this.props.requestSignIn()
    } else {
      this.props.setIsUserDialogOpen(true)
    }
  }

  handleToRootClick = () => {
    this.props.setIsDrawerOpen(false)
  }

  renderProfileIcon = () => {
    if (this.props.isSignedIn) {
      return (<PersonIcon />)
    } else {
      return (<PersonOutlineIcon />)
    }
  }
}
