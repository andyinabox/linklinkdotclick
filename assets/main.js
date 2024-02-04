import { createLink } from './api'
import { handleError } from './errors'
import { Link } from './link'

const linksContainerEl = document.getElementById('links')

// login/signup

const btnLogin = document.querySelector('button[name="login"]')
btnLogin.addEventListener('click', () => {
  const email = prompt('Enter your email:')
  if (email === '') {
    return
  }
  console.log('login user ' + email)
})

// re-ordering links
function sortLinks() {
  const links = linksContainerEl.querySelectorAll('linky-link')
  for (var i = 1; i <= links.length; i++) {
    const l1 = links[i - 1]
    const l2 = links[0]
    const d1 = new Date(l1.data.lastClicked)
    const d2 = new Date(l2.data.lastClicked)

    // swap links
    if (d2 > d1) {
      linksContainerEl.replaceChild(l2, l1)
      linksContainerEl.insertBefore(l1, l2)
    }
  }
}
document.addEventListener('sort-links', sortLinks)

// adding links
async function handleCreateLink() {
  try {
    const url = prompt('Enter a website or feed URL')

    if (!url) return

    document.body.classList.add('loading')
    const link = await createLink(url)
    Link.create(linksContainerEl, link)

    sortLinks()
  } catch (err) {
    handleError(err)
  } finally {
    document.body.classList.remove('loading')
  }
}
const btnAdd = document.querySelector('button[name="add"]')
btnAdd.addEventListener('click', (e) => handleCreateLink(e))

// edit mode
let editing = false
const btnEdit = document.querySelector('button[name="edit"]')
const editingOffText = btnEdit.innerHTML
btnEdit.addEventListener('click', (e) => {
  if (editing) {
    document.body.classList.remove('editing')
    btnEdit.textContent = editingOffText
    editing = false
  } else {
    document.body.classList.add('editing')
    btnEdit.textContent = 'Done'
    editing = true
  }
})
