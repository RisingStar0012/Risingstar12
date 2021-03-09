import axios from "axios";
import {
  User,
  PrincipalId,
  Principal,
  PrincipalState,
  PrincipalStatus,
  ResourceObject,
} from "../../types";

function convert(user: ResourceObject): Principal {
  return {
    id: user.id,
    status: user.attributes.status as PrincipalStatus,
    name: user.attributes.name as string,
    email: user.attributes.email as string,
  };
}

const state: () => PrincipalState = () => ({
  principalList: [],
});

const getters = {
  principalList: (state: PrincipalState) => (): Principal[] => {
    return state.principalList;
  },

  principalById: (state: PrincipalState) => (
    principalId: PrincipalId
  ): Principal => {
    if (!principalId) {
      return {
        id: "-1",
        status: "UNKNOWN",
        name: "<<ID Missing>>",
        email: "",
      };
    }

    for (const principal of state.principalList) {
      if (principal.id == principalId) {
        return principal;
      }
    }

    // Return id as the name if no matching is found.
    return {
      id: principalId,
      status: "UNKNOWN",
      name: principalId,
      email: "",
    };
  },
};

const actions = {
  async fetchPrincipalList({ commit }: any) {
    const userList: ResourceObject[] = (await axios.get(`/api/user`)).data.data;

    const principalList = userList.map((user) => {
      return convert(user);
    });
    commit("setPrincipalList", principalList);

    return userList;
  },

  async fetchPrincipalById({ commit }: any, principalId: PrincipalId) {
    const principal = convert(
      (await axios.get(`/api/user/${principalId}`)).data.data
    );

    commit("replacePrincipalInList", principal);

    return principal;
  },
};

const mutations = {
  setPrincipalList(state: PrincipalState, principalList: Principal[]) {
    state.principalList = principalList;
  },

  replacePrincipalInList(state: PrincipalState, updatedPrincipal: Principal) {
    const i = state.principalList.findIndex(
      (item: Principal) => item.id == updatedPrincipal.id
    );
    if (i != -1) {
      state.principalList[i] = updatedPrincipal;
    }
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
