import _ from 'lodash'
import PropTypes from 'prop-types'
import React from 'react'
import { Scrollbars as S } from 'react-custom-scrollbars'
import { Link } from 'react-router-dom'
import styled from 'styled-components'
import {
  Drawer as D,
  List,
  ListItem,
  ListItemText,
  SwipeableDrawer as SD
} from '@material-ui/core'
import { dimensions, suras } from '@app/constants'

const Drawer = styled(D)`
  .MuiPaper-root {
    box-sizing: border-box;
    padding-top: ${dimensions.APP_BAR_HEIGHT}px;
  }
`

const Scrollbars = styled(S)`
  height: 100%;
  min-width: ${dimensions.LENGTH_200}px;
`

const SwipeableDrawer = styled(SD)`
  .MuiPaper-root {
    padding-top: ${dimensions.APP_BAR_HEIGHT}px;
  }
`

export default class Component extends React.Component {
  static propTypes = {
    isDrawerOpen: PropTypes.bool.isRequired,
    isMobile: PropTypes.bool.isRequired,
    setIsDrawerOpen: PropTypes.func.isRequired
  }

  render () {
    const renderedSuras = this.renderSuras()
    if (this.props.isMobile) {
      return (
        <SwipeableDrawer
          onClose={this.handleClose}
          onOpen={this.handleOpen}
          open={this.props.isDrawerOpen}
        >
          {renderedSuras}
        </SwipeableDrawer>
      )
    } else {
      return (
        <Drawer variant='permanent'>{renderedSuras}</Drawer>
      )
    }
  }

  handleClose () {
    this.props.setIsDrawerOpen(false)
  }

  handleLinkClick () {
    this.props.setIsDrawerOpen(false)
  }

  handleOpen () {
    this.props.setIsDrawerOpen(true)
  }

  renderSuras () {
    const links = _.keys(suras)
      .sort((a, b) => a - b)
      .map((suraNumber) => {
        return (
          <Link
            key={suraNumber}
            onClick={this.handleLinkClick}
            to={`/sura/${suraNumber}`}
          >
            <ListItem button>
              <ListItemText primary={_.get(suras, [suraNumber, 'suraName'])} />
            </ListItem>
          </Link>
        )
      })
    return (
      <Scrollbars>
        <List>{links}</List>
      </Scrollbars>
    )
  }
}
