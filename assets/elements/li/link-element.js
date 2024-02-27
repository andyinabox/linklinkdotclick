import { eventsMixin, slotsMixin } from '../../lib/mixins'
import { updateLink, deleteLink, patchLink, getLink } from '../../lib/api'
import { handleError } from '../../lib/errors'
class LinkElement extends HTMLLIElement {
  constructor() {
    super()
    this.registerSlots()

    // replace link href
    const href = this.slots.link.getAttribute('data-href')
    if (href) {
      this.siteUrl = href
      this.removeAttribute('data-href')
    }
  }

  async fetchData() {
    try {
      this.loading = true
      const link = await getLink(this.linkId)
      this.data = link
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }

  get data() {
    return { ...this._data }
  }
  set data(d) {
    this._data = d
    this.render()
  }

  set loading(v) {
    if (v) {
      this.classList.add('loading')
      this.broadcast('loading-start')
      this.slots.form.disabled = true
    } else {
      this.classList.remove('loading')
      this.broadcast('loading-stop')
      this.slots.form.disabled = false
    }
  }

  get linkId() {
    return this.slots.form.resourceId
  }

  set linkId(v) {
    this.slots.form.resourceId = v
  }

  render() {
    const data = this.data

    // set form data
    this.slots.form.state = data

    if (!data) return
    const { id, siteName, unreadCount, siteUrl, hideUnreadCount, feedUrl } =
      data

    if (hideUnreadCount) {
      this.slots.count.textContent = ''
    } else if (unreadCount) {
      this.slots.count.textContent = data.unreadCount
    } else if (!feedUrl) {
      this.slots.count.textContent = '?'
    } else {
      this.slots.count.textContent = ''
    }

    this.linkId = id
    this.slots.link.innerText = siteName
    this.slots.link.href = siteUrl
  }

  async onClick() {
    try {
      this.loading = true
      const link = await patchLink(this.linkId, { lastClicked: new Date() })
      this.data = link
      this.broadcast('link-click-success')
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }
  async onUpdate(evt) {
    const state = evt.detail.state
    try {
      this.loading = true
      const link = await updateLink(Object.assign(this.data, state))
      this.data = link
      this.broadcast('link-update-success')
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }
  async onDelete(evt) {
    const state = evt.detail.state
    if (!confirm(`Are you sure you want to delete ${state.siteName}?`)) return
    try {
      this.loading = true
      await deleteLink(state.id)
      this.remove()
    } catch (err) {
      handleError(err)
    } finally {
      this.loading = false
    }
  }

  connectedCallback() {
    this.listen(this.slots.link, 'click', this.onClick)
    this.listen(this.slots.form, 'link-delete-request', this.onDelete)
    this.listen(this.slots.form, 'link-update-request', this.onUpdate)
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}
Object.assign(LinkElement.prototype, eventsMixin, slotsMixin)
customElements.define('link-element', LinkElement, { extends: 'li' })
