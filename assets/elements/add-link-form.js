import { eventsMixin } from '../lib/mixins'
import { createLink } from '../lib/api'
import { ProgressiveForm } from '../lib/progressive-form'
class AddLinkForm extends ProgressiveForm {
  constructor() {
    super()
    this.querySelector('input[type="text"]').remove()
  }

  async handleCreateLink() {
    let url = prompt('Enter a website or feed URL')
    if (!url) return

    try {
      this.broadcast('loading-start')
      const link = await createLink(url)
      console.log('created link', link)
      this.broadcast('link-add-success', { link })
    } catch (err) {
      handleError(err)
    } finally {
      this.broadcast('loading-stop')
    }
  }

  connectedCallback() {
    const submit = this.querySelector('input[type="submit"]')
    this.listen(submit, 'click', this.handleCreateLink)
  }
  disconnectedCallback() {
    this.unlistenAll()
  }
}
Object.assign(AddLinkForm.prototype, eventsMixin)
customElements.define('add-link-form', AddLinkForm, { extends: 'form' })
