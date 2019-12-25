import { connect } from 'react-redux'

import { app as appSlice, user as userSlice } from '@app/slices'
import Component from './Component'

export default connect(
  (state) => {
    return {
      open: state.app.isUserDialogOpen
    }
  },
  (dispatch) => {
    return {
      setIsUserDialogOpen: (isOpen) => dispatch(appSlice.actions.setIsUserDialogOpen(isOpen)),
      userReset: () => dispatch(userSlice.actions.reset())
    }
  }
)(Component)
