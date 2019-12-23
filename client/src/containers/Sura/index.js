import { connect } from 'react-redux'
import { withRouter } from 'react-router'
import { compose } from 'redux'

import { quran } from '@app/slices'
import Component from './Component'

export default compose(
  withRouter,
  connect(
    (state) => {
      const quran = state.quran
      return {
        quran
      }
    },
    (dispatch, ownProps) => {
      return {
        requestSura: () => dispatch(quran.actions.requestSura(ownProps.match.params.number))
      }
    }
  )
)(Component)
