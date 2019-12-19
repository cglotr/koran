import _ from 'lodash'
import PropTypes from 'prop-types'
import React from 'react'
import { Column, Verse } from '@app/components'
import { Typography } from '@material-ui/core'

export default class Component extends React.Component {
  static propTypes = {
    quran: PropTypes.object.isRequired,
    requestSura: PropTypes.func.isRequired,
    match: PropTypes.shape({
      params: PropTypes.shape({
        number: PropTypes.string.isRequired
      })
    })
  }

  componentDidMount () {
    this.props.requestSura()
  }

  componentDidUpdate (prevProps) {
    const prevSuraNumber = _.get(prevProps, ['match', 'params', 'number'])
    const suraNumber = _.get(this.props, ['match', 'params', 'number'])
    if (suraNumber !== prevSuraNumber) {
      this.props.requestSura()
    }
  }

  render () {
    const suraNumber = this.props.match.params.number
    const sura = _.get(this.props.quran, suraNumber)
    const verses = _.keys(sura).sort((a, b) => parseInt(a) - parseInt(b))
    return (
      <Column>
        <Typography variant='h5'>Sura {suraNumber}</Typography>
        <Column>
          {
            verses.map((verse) => {
              const ayah = _.get(sura, [verse, 'ayah'])
              const translation = _.get(sura, [verse, 'translation'])
              return (
                <Verse
                  ayah={ayah}
                  key={verse}
                  translation={translation}
                  verseNumber={verse}
                />
              )
            })
          }
        </Column>
      </Column>
    )
  }
}
