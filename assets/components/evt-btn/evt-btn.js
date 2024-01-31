import { Component } from '../component'
import './evt-btn.css'

const DATA_ATTRIBUTE_PREFIX = 'data-e-'

export class EvtBtn extends Component {
  constructor() {
    super()
    this.eventName = this.getAttribute('data-event-name')
    if (!this.eventName) {
      throw new Error('EvtBtn: no event name set')
    }
    this.data = {}
    for (const attr of this.attributes) {
      if (attr.name.indexOf(DATA_ATTRIBUTE_PREFIX) === 0) {
        const key = attr.name.replace(DATA_ATTRIBUTE_PREFIX, '')
        this.data[key] = attr.value // keep in mind this will always be a string
      }
    }
  }
  connectedCallback() {
    this.addEventListener('click', (event) => {
      this.dispatchEvent(
        new CustomEvent(name, {
          bubbles: true,
          detail: this.data,
        })
      )
    })
  }
}
customElements.define('linky-evt-btn', EvtBtn)
