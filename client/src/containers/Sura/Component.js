import _ from 'lodash'
import PropTypes from 'prop-types'
import React from 'react'
import styled from 'styled-components'
import { Column, Footer, Verse } from '@app/components'
import { Typography } from '@material-ui/core'

import { dimensions, suras } from '@app/constants'

const PaddedColumn = styled(Column)`
  padding-bottom: ${dimensions.PADDING_LARGE}px;
  padding-top: ${dimensions.PADDING_LARGE}px;
`

export default class Component extends React.Component {
  static propTypes = {
    isUserSignedIn: PropTypes.bool.isRequired,
    match: PropTypes.shape({
      params: PropTypes.shape({
        number: PropTypes.string.isRequired
      })
    }),
    quran: PropTypes.object.isRequired,
    read: PropTypes.object,
    requestSura: PropTypes.func.isRequired,
    setSuraVerseRead: PropTypes.func.isRequired
  }

  componentDidMount () {
    this.props.requestSura()
  }

  componentDidUpdate (prevProps) {
    const prevSuraNumber = _.get(prevProps, ['match', 'params', 'number'])
    const suraNumber = _.get(this.props, ['match', 'params', 'number'])
    if (suraNumber !== prevSuraNumber || this.props.isUserSignedIn !== prevProps.isUserSignedIn) {
      this.props.requestSura()
    }
  }

  render () {
    const suraNumber = this.props.match.params.number
    const sura = _.get(this.props.quran, suraNumber)
    const verses = _.keys(sura).sort((a, b) => parseInt(a) - parseInt(b))
    const suraName = _.get(suras, [suraNumber, 'suraName'])
    const suraNameTranslation = _.get(suras, [suraNumber, 'suraNameTranslation'])
    return (
      <Column>
        <PaddedColumn>
          <Typography align='center' variant='h5'>{suraName}</Typography>
          <Typography align='center' variant='subtitle1'>{suraNameTranslation}</Typography>
        </PaddedColumn>
        {
          verses.map((verse) => {
            const ayah = _.get(sura, [verse, 'ayah'])
            const isRead = _.get(this.props.read, [suraNumber, verse], false)
            const translation = _.get(sura, [verse, 'translation'])
            return (
              <Verse
                ayah={ayah}
                isCheckboxEnabled={this.props.isUserSignedIn}
                isRead={isRead}
                key={verse}
                onCheckboxChange={this.handleCheckboxChange(verse)}
                translation={translation}
                verseNumber={verse}
              />
            )
          })
        }
        <Footer />
      </Column>
    )
  }

  handleCheckboxChange = (verseNumber) => (isRead) => {
    const suraNumber = this.props.match.params.number
    this.props.setSuraVerseRead(suraNumber, verseNumber, isRead)
  }
}
