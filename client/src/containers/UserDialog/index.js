import { connect } from 'react-redux'

import { user as userSlice } from '@app/slices'
import Component from './Component'

export default connect(
  (state) => {
    return {
      open: state.app.isUserDialogOpen
    }
  },
  (dispatch) => {
    return {
      requestSignOut: () => dispatch(userSlice.actions.requestSignOut())
    }
  }
)(Component)
