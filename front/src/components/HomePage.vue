<template>
  <div>
    <!-- Formulário de cadastro de doador -->
    <div role="form">
      <h2>Cadastro de Doador</h2>
      <input type="text" v-model="nome" placeholder="Nome"><br>
      <input type="text" v-model="cpf" placeholder="CPF"><br>
      <input type="text" v-model="contato" placeholder="Contato"><br>
      <select v-model="tipoSanguineo">
        <option disabled value="">Selecione o tipo sanguíneo</option>
        <option value="A">A</option>
        <option value="B">B</option>
        <option value="AB">AB</option>
        <option value="O">O</option>
      </select><br>
      <select v-model="fatorRh">
        <option disabled value="">Fator RH</option>
        <option value="POSITIVO">Positivo</option>
        <option value="NEGATIVO">Negativo</option>
      </select><br>
      <button @click="adicionarDoador">Adicionar Doador</button>
    </div>

    <!-- Tabela de doadores -->
    <div>
      <h2>Lista de Doadores</h2>
      <button @click="acaoSelecionada(doador, 'buscar')">Buscar</button>
      <br>
      <table border="1">
        <thead>
          <tr>
            <th>Código</th>
            <th>Nome</th>
            <th>CPF</th>
            <th>Contato</th>
            <th>Tipo Sanguíneo</th>
            <th>RH</th>
            <th>Ações</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="doador in doadores" :key="doador.codigo">
            <td>{{ doador.codigo  }}</td>
            <td>{{ doador.nome }}</td>
            <td>{{ doador.cpf }}</td>
            <td>{{ doador.contato }}</td>
            <td>{{ doador.tipoSanguineo }}</td>
            <td>{{ doador.fatorRh }}</td>
            <td>
              <button @click="acaoSelecionada(doador, 'editar')">Editar</button>
              <button @click="acaoSelecionada(doador, 'inativar')">Remover</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Modal de Edição -->
    <div v-if="mostrarDropdownEditar" class="modal">
      <div class="modal-content">
      <span class="close" @click="fecharDropdown">&times;</span>
        <h3>Editar Doador</h3>
        <input type="text" v-model="doadorEditado.nome"  placeholder="Nome"><br>
        <input type="text" v-model="doadorEditado.cpf"  placeholder="CPF"><br>
        <input type="text" v-model="doadorEditado.contato"  placeholder="Contato"><br>
        <select v-model="doadorEditado.tipoSanguineo"  placeholder="Tipo Sanguineo">
          <option value="A">A</option>
          <option value="B">B</option>
          <option value="AB">AB</option>
          <option value="O">O</option>
        </select><br>
        <select v-model="doadorEditado.fatorRh"  placeholder="Fator RH">
          <option value="POSITIVO">Positivo</option>
          <option value="NEGATIVO">Negativo</option>
        </select><br>
        <button @click="editarDoador">Salvar</button>
      </div>
    </div>

    <!-- Popup de Inativação -->
    <div v-if="mostrarPopupInativar" class="popup-inativar">
      <div class="popup-content">
        <span class="close" @click="fecharDropdown">&times;</span>
        <h3>Confirmar Inativação</h3>
        <p>Deseja realmente inativar o doador {{ doadorInativar.nome }}?</p>
        <button @click="inativarDoador">Sim</button>
        <button @click="fecharPopupInativar">Não</button>
      </div>
    </div>
  </div>


  <div v-if="mostrarDropdownBuscar" class="modal">
      <div class="modal-content">
      <span class="close" @click="fecharDropdown">&times;</span>
        <h3>Buscar Doador</h3>
        <input type="text" v-model="doadorBuscavel.nome" placeholder="Nome"><br>
        <input type="text" v-model="doadorBuscavel.cpf"  placeholder="CPF"><br>
        <input type="text" v-model="doadorBuscavel.contato" placeholder="Contato"><br>
        <select v-model="doadorBuscavel.tipoSanguineo">
          <option disabled value="">Tipo Sanguineo</option>
          <option value="A">A</option>
          <option value="B">B</option>
          <option value="AB">AB</option>
          <option value="O">O</option>
          <option value="">TODOS</option>
        </select><br>
        <select v-model="doadorBuscavel.fatorRh">
          <option disabled value="">Fator RH</option>
          <option value="POSITIVO">Positivo</option>
          <option value="NEGATIVO">Negativo</option>
          <option value="">TODOS</option>
        </select><br>
        <button @click="buscarTodosDoadores">Salvar</button>
      </div>
    </div>
  
</template>


<script>
import api from '../services';
export default {
  data() {
    return {
      nome: '',
      cpf: '',
      contato: '',
      tipoSanguineo: '',
      fatorRh: '',
      filtro: '',
      doadores: [],
      mostrarDropdownEditar: false,
      doadorEditado: {
        codigo:'',
        nome: '',
        cpf: '',
        contato: '',
        tipoSanguineo: '',
        fatorRh: ''
      },
      mostrarPopupInativar: false,
      doadorInativar: {},
      mostrarDropdownBuscar: false,
      doadorBuscavel: {
        codigo: -1,
        nome: '',
        cpf: '',
        contato: '',
        tipoSanguineo: '',
        fatorRh: ''
      },
    };
  },
  mounted(){
    this.buscarTodosDoadores()
  },

  methods: {
      async adicionarDoador() {
      // Simulação de adição de doador
      // this.doadores.push({
      //   nome: this.nome,
      //   cpf: this.cpf,
      //   contato: this.contato,
      //   tipoSanguineo: this.tipoSanguineo,
      //   fatorRh: this.fatorRh
      // });
      //api.get("/ping")
    //  const response = await api.post(api.defaults.baseURL + 'dados', this.doadores);
     console.log(this.tipoSanguineo);
      await api.post('dados', {
      nome: this.nome,
      cpf: this.cpf,
      contato: this.contato,
      tipoSanguineo: this.tipoSanguineo,
      fatorRh: this.fatorRh,
      situacao: "Ativo"
}).then(response => {
  this.buscarTodosDoadores()
  console.log(response);
});

  // Limpe os campos do formulário após a adição
      this.nome = '';
      this.cpf = '';
      this.contato = '';
      this.tipoSanguineo = '';
      this.fatorRh = '';

    },
    acaoSelecionada(doador, acao) {
      if (acao === 'editar') {
        this.mostrarDropdownEditar = true;
        this.doadorEditado = Object.assign({}, doador); // Clonar o objeto para evitar mutações indesejadas
      } else if (acao === 'inativar') {
        this.mostrarPopupInativar = true;
        this.doadorInativar = doador;
      } else if (acao == 'buscar') {
        this.mostrarDropdownBuscar = true;
        this.LimpaDepois()

      }
    },
    fecharPopupInativar() {
      this.mostrarPopupInativar = false;
    },

    fecharDropdown() {
    this.mostrarDropdownEditar = false; 
    if(this.mostrarDropdownBuscar){
      this.mostrarDropdownBuscar = false;
      this.LimpaDepois()
    }
    this.mostrarPopupInativar = false;
  },


  LimpaDepois() {
       this.doadorBuscavel.codigo = - 1   
       this.doadorBuscavel.nome = ''
       this.doadorBuscavel.cpf = ''
       this.doadorBuscavel.contato = ''
       this.doadorBuscavel.tipoSanguineo = ''
       this.doadorBuscavel.fatorRh = ''
  },

    
  async editarDoador() {
      await api.put('dados/'+this.doadorEditado.codigo, {
      nome: this.doadorEditado.nome,
      cpf: this.doadorEditado.cpf,
      contato: this.doadorEditado.contato,
      tipoSanguineo: this.doadorEditado.tipoSanguineo,
      fatorRh: this.doadorEditado.fatorRh
    
})
      location.reload();

    },

   async inativarDoador() {
      await api.delete('/dados/'+this.doadorInativar.codigo) 
      location.reload()
    },
    async buscarTodosDoadores() {
      await api.post('informacoes',this.doadorBuscavel).then(response => {
        this.doadores = []
        response.data?.map(item => this.doadores.push(item))
        
      })
      this.mostrarDropdownBuscar = false;
    },
  }
};

</script>

<style>
body {
  font-family: Arial, sans-serif;
  margin: 0;
  padding: 0;
  background-color: #f0f0f0;
}

div[role="form"] {
  max-width: 600px;
  margin: 20px auto;
  padding: 20px;
  background-color: #fff;
  box-shadow: 0 0 10px rgba(0,0,0,0.1);
}

table {
  width: 100%;
  border-collapse: collapse;
}

table, th, td {
  border: 1px solid #ddd;
}

th, td {
  padding: 8px;
  text-align: left;
}

.modal,.popup-inativar {
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgba(0,0,0,0.4);
}

.modal-content,.popup-content {
  position: relative;
  background-color: #fefefe;
  margin: 15% auto;
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
  box-shadow: 0 4px 8px rgba(0,0,0,0.2);
}

button {
  background-color: #4CAF50;
  color: white;
  padding: 10px 20px;
  border: none;
  cursor: pointer;
  margin: 5px;
  width: 100%;
}

button:hover {
  background-color: #45a049;
}

input, select {
  width: 100%;
  padding: 12px 20px;
  margin: 8px 0;
  display: inline-block;
  border: 1px solid #ccc;
  box-sizing: border-box;
}

h2 {
  text-align: center;
  margin-bottom: 20px;
}

.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5); /* Fundo semi-transparente para destacar o modal */
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal-content {
  position: relative;
  background-color: white;
  padding: 20px;
}

.close {
  position: absolute;
  top: 10px;
  right: 10px;
  font-size: 24px;
  color: #000000;
  cursor: pointer;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}

</style>
