import { connect } from 'react-redux'
import { app } from '@app/slices'
import Component from './Component'

export default connect(
  (state) => {
    const isDrawerOpen = state.app.isDrawerOpen
    const userEmail = state.user.email
    return {
      isDrawerOpen,
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
