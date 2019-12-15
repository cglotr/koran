import _ from 'lodash'
import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import D from '@material-ui/core/Drawer'
import List from '@material-ui/core/List'
import ListItem from '@material-ui/core/ListItem'
import ListItemText from '@material-ui/core/ListItemText'
import { suras } from '@app/constants'

export default class Drawer extends Component {
  render () {
    return (
      <>
        <D variant='permanent'>
          <List>{this.renderSuras()}</List>
        </D>
      </>
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
