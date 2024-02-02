import { getLink, updateLink, deleteLink } from '../../lib/api'
import { handleError } from '../../lib/errors'
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

  render() {
    const data = this.data

    this.slots.link.href = data.siteUrl
    this.slots.link.textContent = data.siteName

    if (data.unreadCount) {
      this.slots.count.textContent = `(${data.unreadCount})`
    }

    this.setAttribute('data-id', data.id)
  }

  async fetchData(id) {
    this.classList.add('loading')
    try {
      this.data = await getLink(id)
      this.dispatchEvent(new CustomEvent('sort-links', { bubbles: true }))
    } catch (err) {
      handleError(err)
    } finally {
      this.classList.remove('loading')
    }
  }

  async handleClick(e) {
    const link = this.data
    link.lastClicked = new Date().toJSON()
    try {
      const updatedLink = await updateLink(link)
      this.data = updatedLink
      this.dispatchEvent(new CustomEvent('sort-links', { bubbles: true }))
    } catch (err) {
      handleError(err)
    }
  }

  async handleDelete(e) {
    const { id, siteName } = this.data

    if (!confirm(`Delete link "${siteName}"?`)) return

    try {
      await deleteLink(id)
      this.remove()
    } catch (err) {
      handleError(err)
    }
  }

  async handleRename(e) {
    const data = this.data
    const name = prompt(`Enter a new name for "${data.siteName}"`)
    if (!name) return
    data.siteName = name
    try {
      const updated = await updateLink(data)
      this.data = updated
    } catch (err) {
      handleError(err)
    }
  }

  connectedCallback() {
    this.onClick = (e) => this.handleClick(e)
    this.slots.link.addEventListener('click', this.onClick)

    const edit = this.slots.edit
    this.onDelete = (e) => this.handleDelete(e)
    edit
      .querySelector('[name="delete"]')
      .addEventListener('click', this.onDelete)
    this.onRename = (e) => this.handleRename(e)
    edit
      .querySelector('[name="rename"]')
      .addEventListener('click', this.onRename)
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
