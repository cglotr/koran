import _ from 'lodash'
import React, { Component } from 'react'
import { Scrollbars } from 'react-custom-scrollbars'
import { Link } from 'react-router-dom'
import styled from 'styled-components'
import D from '@material-ui/core/Drawer'
import List from '@material-ui/core/List'
import ListItem from '@material-ui/core/ListItem'
import ListItemText from '@material-ui/core/ListItemText'
import { dimensions, suras } from '@app/constants'

const DD = styled(D)`
  .MuiDrawer-paper {
    box-sizing: border-box;
    width: ${dimensions.LENGTH_200}px;
  }
`

const S = styled(Scrollbars)`
  height: 100%;
  width: 100%;
`

export default class Drawer extends Component {
  render () {
    return (
      <DD variant='permanent'>
        <S>
          <List>{this.renderSuras()}</List>
        </S>
      </DD>
    )
  }

  renderSuras () {
    return _.keys(suras)
      .sort((a, b) => a - b)
      .map((suraNumber) => {
        return (
          <Link key={suraNumber} to={`/sura/${suraNumber}`}>
            <ListItem button>
              <ListItemText primary={_.get(suras, [suraNumber, 'suraName'])} />
            </ListItem>
          </Link>
        )
      })
  }
}
