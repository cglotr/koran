import PropTypes from 'prop-types'
import React from 'react'

class Component extends React.Component {
  componentDidMount () {
    this.props.fetchSura()
  }

  render () {
    const suraNumber = this.props.match.params.number
    return (
      <>
        <h1>Sura {suraNumber}</h1>
      </>
    )
  }
}

Component.propTypes = {
  fetchSura: PropTypes.func.isRequired,
  match: PropTypes.shape({
    params: PropTypes.shape({
      number: PropTypes.number.isRequired
    })
  })
}

export default Component
