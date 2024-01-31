import { slotsMixin } from '../mixins'
import './link.css'
class Link extends HTMLElement {
  constructor() {
    super()
    this.registerSlots()
  }
  async getData() {
    const id = this.getAttribute('data-id')
    const response = await fetch(`/api/links/${id}?refresh`)
    const json = await response.json()
    this.data = json
    this.render()
  }
  async deleteLink() {}
  async updateLink() {}
  render() {
    this.slots.count.innerHTML = `(${this.data.unreadCount})`
  }
  connectedCallback() {
    this.getData()
    this.slots.delete.addEventListener('click', (event) => {})
  }
}

Object.assign(Link.prototype, slotsMixin)

customElements.define('linky-link', Link)
