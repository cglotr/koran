import PropTypes from 'prop-types'
import React from 'react'
import {
  Dialog,
  List,
  ListItem,
  ListItemText
} from '@material-ui/core'

export default class Component extends React.Component {
  static propTypes = {
    open: PropTypes.bool.isRequired,
    requestSignOut: PropTypes.func.isRequired,
    setIsUserDialogOpen: PropTypes.func.isRequired
  }

  render () {
    return (
      <Dialog
        fullWidth={true}
        onBackdropClick={this.handleBackdropClick}
        open={this.props.open}
      >
        <List>
          <ListItem button onClick={this.handleSignOutClick}>
            <ListItemText primary="Sign out" />
          </ListItem>
        </List>
      </Dialog>
    )
  }

  handleBackdropClick = () => {
    this.props.setIsUserDialogOpen(false)
  }

  handleSignOutClick = () => {
    this.props.requestSignOut()
  }
}
