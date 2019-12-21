import { connect } from 'react-redux'
import { app } from '@app/slices'
import Component from './Component'

export default connect(
  (state) => {
    const isDrawerOpen = state.app.isDrawerOpen
    return {
      isDrawerOpen
    }
  },
  (dispatch) => {
    return {
      setIsDrawerOpen: (open) => dispatch(app.actions.setIsDrawerOpen(open))
    }
  }
)(Component)
