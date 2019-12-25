import { connect } from 'react-redux'

import { user as userSelector } from '@app/selectors'
import { app } from '@app/slices'
import Component from './Component'

export default connect(
  (state) => {
    const isDrawerOpen = state.app.isDrawerOpen
    const userEmail = state.user.email
    return {
      isDrawerOpen,
      isSignedIn: userSelector.isSignedIn(state),
      userEmail
    }
  },
  (dispatch) => {
    return {
      requestSignIn: () => dispatch(app.actions.requestSignIn()),
      setIsDrawerOpen: (open) => dispatch(app.actions.setIsDrawerOpen(open))
    }
  }
)(Component)
