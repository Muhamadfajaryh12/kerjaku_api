<h6>Menjalankan project dengan</h6>
<span style="color:red">Merah</span>, <span style="color:blue">Biru</span>
  air
</div>
<h1>Dokumentasi API</h1>
<table>
  <thead>
    <tr>
      <th>Endpoint</th>
      <td>Method</td>
      <th>Deskripsi</th>
      <th>Authorization</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td >http://127.0.0.1:3000/api/register</td>
      <td >POST</td>
      <td>Endpoint untuk mendaftarkan user</td>
      <td>-</td>
    </tr>
     <tr>
      <td >http://127.0.0.1:3000/api/login</td>
      <td >POST</td>
      <td>Endpoint untuk user login</td>
      <td>-</td>
    </tr>
     <tr>
      <td >http://127.0.0.1:3000/api/profile/:id</td>
      <td >GET</td>
      <td>Endpoint untuk melihat data profile </td>
      <td>Bearer Token</td>
    </tr>
      <tr>
      <td >http://127.0.0.1:3000/api/profile</td>
      <td >POST</td>
      <td>Endpoint untuk menginput data profile </td>
      <td>Bearer Token</td>
    </tr>
    <tr>
      <td >http://127.0.0.1:3000/api/profile/:id</td>
      <td >PUT</td>
      <td>Endpoint untuk mengupdate data profile </td>
      <td>Bearer Token</td>
    </tr>
      <tr>
      <td >http://127.0.0.1:3000/api/experience</td>
      <td >POST</td>
      <td>Endpoint untuk menginput data experience </td>
      <td>Bearer Token</td>
    </tr>
    <tr>
      <td >http://127.0.0.1:3000/api/experience/:id</td>
      <td >PUT</td>
      <td>Endpoint untuk mengupdate data experience </td>
      <td>Bearer Token</td>
    </tr>
          <tr>
      <td >http://127.0.0.1:3000/api/experience/:id</td>
      <td >DELETE</td>
      <td>Endpoint untuk menghapus data experience </td>
      <td>Bearer Token</td>
    </tr>
          <tr>
      <td >http://127.0.0.1:3000/api/company</td>
      <td >POST</td>
      <td>Endpoint untuk menginput data company </td>
      <td>Bearer Token</td>
    </tr>
          <tr>
      <td >http://127.0.0.1:3000/api/company/:id</td>
      <td >PUT</td>
      <td>Endpoint untuk mengubah data company </td>
      <td>Bearer Token</td>
    </tr>
     <tr>
      <td >http://127.0.0.1:3000/api/company?search=&type=&category</td>
      <td >GET</td>
      <td>Endpoint untuk mendapatkan  data company </td>
      <td>-</td>
    </tr>
     <tr>
      <td >http://127.0.0.1:3000/api/company/:id</td>
      <td >GET</td>
      <td>Endpoint untuk mendapatkan detail data company </td>
      <td>-</td>
    </tr>
     <tr>
      <td >http://127.0.0.1:3000/api/vacancy</td>
      <td >POST</td>
      <td>Endpoint untuk menginput data vacancy </td>
      <td>Bearer Token</td>
    </tr>
       <tr>
      <td >http://127.0.0.1:3000/api/vacancy/:id</td>
      <td >PUT</td>
      <td>Endpoint untuk mengubah data vacancy </td>
      <td>Bearer Token</td>
    </tr>
       <tr>
      <td >http://127.0.0.1:3000/api/vacancy/:id</td>
      <td >GET</td>
      <td>Endpoint untuk mendapat data detail vacancy </td>
      <td>-</td>
    </tr>
        <tr>
      <td >http://127.0.0.1:3000/api/vacancy?search=&category=&location=</td>
      <td >GET</td>
      <td>Endpoint untuk mendapat data vacancy </td>
      <td>-</td>
    </tr>
        <tr>
      <td >http://127.0.0.1:3000/api/application</td>
      <td >POST</td>
      <td>Endpoint untuk menginput data application </td>
      <td>Bearer Token</td>
    </tr>
       <tr>
      <td >http://127.0.0.1:3000/api/application/:id</td>
      <td >PUT</td>
      <td>Endpoint untuk mengubah data application </td>
      <td>Bearer Token</td>
    </tr>
       <tr>
      <td >http://127.0.0.1:3000/api/application/:id</td>
      <td >DELETE</td>
      <td>Endpoint untuk menghapus data application </td>
      <td>Bearer Token</td>
    </tr>
       <tr>
      <td >http://127.0.0.1:3000/api/application?search=&user=</td>
      <td >GET</td>
      <td>Endpoint untuk mendapatkan data application </td>
      <td>-</td>
    </tr>
      <tr>
      <td >http://127.0.0.1:3000/api/application?search=&vacancy=</td>
      <td >GET</td>
      <td>Endpoint untuk mendapatkan data application </td>
      <td>Bearer Token</td>
    </tr>
  </tbody>
  
</table>
