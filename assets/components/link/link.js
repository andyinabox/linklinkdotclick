import { slotsMixin } from '../../lib/mixins'
import { getLink } from '../../lib/api'

import './link.css'
class Link extends HTMLLIElement {
  constructor() {
    super()
    this.registerSlots()
  }
  async getData() {
    const id = this.getAttribute('data-id')
    this.data = await getLink(id)
    this.render()
  }
  render() {
    this.slots.count.innerHTML = `(${this.data.unreadCount})`
  }
  connectedCallback() {
    this.getData()
    this.slots.delete.addEventListener('click', (event) => {})
  }
}

Object.assign(Link.prototype, slotsMixin)

customElements.define('linky-link', Link, { extends: 'li' })
