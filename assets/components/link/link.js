import { getLink, updateLink, deleteLink } from '../../lib/api'
import { Component } from '../component'

import './link.css'
export class Link extends Component {
  constructor() {
    super()
    const id = this.getAttribute('data-id')
    if (id) this.fetchData(id)
  }
  set data(d) {
    this._data = d
    this.render()
  }
  get data() {
    return { ...this._data }
  }
  async fetchData(id) {
    this.classList.add('loading')
    this.data = await getLink(id)
    this.classList.remove('loading')
  }

  render() {
    const data = this.data

    this.slots.link.href = data.siteUrl
    this.slots.link.textContent = data.siteName
    this.slots.count.textContent = data.unreadCount

    this.setAttribute('data-id', data.id)
  }

  async handleClick(e) {
    const link = this.data
    link.lastClicked = new Date().toJSON()
    const updatedLink = await updateLink(link)
    this.data = updatedLink
  }

  async handleDelete(e) {
    const { id, siteName } = this.data

    if (!confirm(`Delete link "${siteName}"?`)) return

    try {
      await deleteLink(id)
      this.remove()
    } catch (e) {
      console.error(e)
    }
  }

  connectedCallback() {
    this.onClick = (e) => this.handleClick(e)
    this.slots.link.addEventListener('click', this.onClick)
    this.onDelete = (e) => this.handleDelete(e)
    this.slots.edit
      .querySelector('[name="delete"]')
      .addEventListener('click', this.onDelete)
  }
  disconnectedCallback() {
    this.slots.link.removeEventListener('click', this.onClick)
    this.slots.edit
      .querySelector('[name="delete"]')
      .removeEventListener('click', this.onDelete)
  }
}

Link.create = function (parentEl, data) {
  const linkTmpl = document.getElementById('tmpl-link')
  const linkEl = linkTmpl.content.firstElementChild.cloneNode(true)
  parentEl.appendChild(linkEl)
  linkEl.data = data
  return linkEl
}

customElements.define('linky-link', Link)
