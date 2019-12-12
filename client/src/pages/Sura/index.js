import { connect } from 'react-redux'
import { withRouter } from 'react-router'
import { compose } from 'redux'
import { quran } from '@app/slices'
import Component from './Component'

const withConnect = connect(
  () => {
    return {}
  },
  (dispatch, ownProps) => {
    return {
      requestSura: () => dispatch(quran.actions.requestSura(ownProps.match.params.number))
    }
  }
)

export default compose(
  withRouter,
  withConnect
)(Component)
