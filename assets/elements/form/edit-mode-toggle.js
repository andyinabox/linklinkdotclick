import { BaseForm } from './base-form'
class EditModeToggle extends BaseForm {
  constructor() {
    super()
    this.btn = this.querySelector('button')
  }
  get isEditing() {
    return this._editing
  }
  set isEditing(v) {
    if (v) {
      this.btn.innerText = 'Done'
      this._editing = true
      this.broadcast('edit-mode-start')
    } else {
      this.btn.innerText = 'Edit'
      this._editing = false
      this.broadcast('edit-mode-stop')
    }
  }
  onSubmit() {
    this.isEditing = !this.isEditing
  }
}
customElements.define('edit-mode-toggle', EditModeToggle, { extends: 'form' })
