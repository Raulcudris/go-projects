import { useEffect, useState } from "react";

function App() {
  const [name , setName] = useState('')
  const [user, setUsers] = useState([])

  useEffect(() => {
    async function loadUsers() {
      const response = await fetch(import.meta.env.VITE_API+"/users");
      const data = await response.json();
      console.log(data)
      setUsers(data.user)
    }
    loadUsers()
  }, [])
  

  const handleSubmit =  async (e) =>{
    e.preventDefault()
    const response = await fetch(import.meta.env.VITE_API+'/users',{
      method: 'POST',
      body: JSON.stringify({name}),
      headers:{
        "Content-Type":"application/json"
      }
    })
    const data = await response.json()
    console.log(data)
  }

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input 
        type="name" 
        placeholder="Coloca tu nombre"
        onChange={(e)=> setName(e.target.value)}/>
        <button>Guardar</button>
      </form>
      <ul>
        {/* { user.map(users =>(
            <li key={users._id}>
              {users.name}
              </li>
         ))
        } */}
      </ul>
    </div>
  )
}

export default App
