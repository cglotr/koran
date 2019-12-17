import { connect } from 'react-redux'
import { app } from '@app/slices'
import Component from './Component'

export default connect(
  (state) => {
    return {
      isDrawerOpen: state.app.isDrawerOpen,
      isMobile: state.app.isMobile
    }
  },
  (dispatch) => {
    return {
      setIsDrawerOpen: (open) => dispatch(app.actions.setIsDrawerOpen(open))
    }
  }
)(Component)
