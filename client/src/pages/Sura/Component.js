import React from 'react'

export default class Component extends React.Component {
  render () {
    return (
      <div>Sura {this.props.match.params.number}</div>
    )
  }
}
