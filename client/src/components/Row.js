import React from 'react'
import styled from 'styled-components'

const D = styled.div`
  display: flex:
  flex-direction: row;
`

export default class Row extends React.Component {
  render () {
    return (
      <D {...this.props}>{this.props.children}</D>
    )
  }
}
