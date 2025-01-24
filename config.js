export default class config {
  static headers() {
    return {
      "content-type": "application/json",
    };
  }

  static getOrchUrl() {
    return "http://localhost:8001/orch";
  }

  static getCoreUrl() {
    return "http://localhost:8002/core";
  }
}
