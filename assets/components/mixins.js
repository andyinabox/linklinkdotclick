export const slotsMixin = {
  slots: {},
  registerSlots() {
    this.querySelectorAll('[slot]').forEach((el) => {
      this.slots[el.getAttribute('slot')] = el
    })
  },
}
