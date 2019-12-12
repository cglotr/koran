import _ from 'lodash'
import PropTypes from 'prop-types'
import React from 'react'

class Component extends React.Component {
  componentDidMount () {
    this.props.requestSura()
  }

  render () {
    const suraNumber = this.props.match.params.number
    const sura = _.get(this.props.quran, suraNumber)
    const verses = _.keys(sura).sort((a, b) => parseInt(a) - parseInt(b))
    return (
      <>
        <h1>Sura {suraNumber}</h1>
        <div
          style={{
            alignItems: 'flex-end',
            display: 'flex',
            flexDirection: 'column'
          }}
        >
          {
            verses.map((verse) => {
              const ayah = _.get(sura, [verse, 'ayah'])
              return (
                <div key={verse}>{`(${verse}) ${ayah}`}</div>
              )
            })
          }
        </div>
      </>
    )
  }
}

Component.propTypes = {
  quran: PropTypes.object.isRequired,
  requestSura: PropTypes.func.isRequired,
  match: PropTypes.shape({
    params: PropTypes.shape({
      number: PropTypes.string.isRequired
    })
  })
}

export default Component
