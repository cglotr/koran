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

const AppBar = styled(AB)`
  &.MuiAppBar-root {
    z-index: 9999;
  }
`

export default class Component extends React.Component {
  static propTypes = {
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
          <Typography noWrap variant='h6'>Quran</Typography>
        </Toolbar>
      </AppBar>
    )
  }

  handleClick = () => {
    this.props.setIsDrawerOpen(true)
  }
}