import PropTypes from 'prop-types'
import React from 'react'
import styled from 'styled-components'
import {
  AppBar as AB,
  IconButton,
  Toolbar,
  Typography
} from '@material-ui/core'
import MenuIcon from '@material-ui/icons/Menu'

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
    setIsDrawerOpen: PropTypes.func.isRequired
  }

  render () {
    return (
      <AppBar position='fixed'>
        <Toolbar>
          <IconButton
            color='inherit'
            edge='start'
            onClick={this.handleClick}
          >
            <MenuIcon />
          </IconButton>
          <Typography noWrap variant='h6'>Koran</Typography>
        </Toolbar>
      </AppBar>
    )
  }

  handleClick = () => {
    this.props.setIsDrawerOpen(!this.props.isDrawerOpen)
  }
}
