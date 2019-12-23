import { connect } from 'react-redux'

import { app as appSlice } from '@app/slices'
import Component from './Component'

export default connect(
  (state) => {
    const isMobile = state.app.isMobile
    const width = state.app.width
    return {
      isMobile,
      width
    }
  },
  (dispatch) => {
    return {
      setWidth: (clientWidth) => dispatch(appSlice.actions.setWidth(clientWidth))
    }
  }
)(Component)
