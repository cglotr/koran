import _ from 'lodash'
import React, { Component } from 'react'
import { quran } from '@app/constants'

export default class Quran extends Component {
  render () {
    return (
      <>
        <h1>Quran</h1>
        <div>
          {
            _.keys(quran)
              .map((key) => {
                return (
                  [parseInt(key), 1, parseInt(quran[key])]
                )
              })
              .sort((a, b) => a.key - b.key)
              .map((sorted) => {
                return (
                  <div>
                    <h2>{sorted[0]}</h2>
                    <div>{sorted[1]} - {sorted[2]}</div>
                  </div>
                )
              })
          }
        </div>
      </>
    )
  }
}
