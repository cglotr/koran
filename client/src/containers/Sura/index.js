import { connect } from 'react-redux'
import { withRouter } from 'react-router'
import { compose } from 'redux'

import { user as userSelector } from '@app/selectors'
import {
  quran,
  read as readSlice
} from '@app/slices'
import Component from './Component'

export default compose(
  withRouter,
  connect(
    (state) => {
      const { quran, read } = state
      const isUserSignedIn = userSelector.isSignedIn(state)
      return {
        isUserSignedIn,
        quran,
        read
      }
    },
    (dispatch, ownProps) => {
      return {
        requestSura: () => dispatch(quran.actions.requestSura(ownProps.match.params.number)),
        setSuraVerseRead: (suraNumber, verseNumber, isRead) => {
          dispatch(readSlice.actions.setSuraVerseRead({ isRead, suraNumber, verseNumber }))
        }
      }
    }
  )
)(Component)
