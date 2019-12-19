import React from 'react'
import styled from 'styled-components'

const D = styled.div`
  align-items: stretch;
  display: flex;
  flex-direction: column;
`

export default class Column extends React.Component {
  render () {
    return (
      <D {...this.props}>{this.props.children}</D>
    )
  }
}
