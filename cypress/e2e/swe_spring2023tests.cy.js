describe("The Home Page", () => {
  it("successfully loads", () => {
    cy.visit("http://localhost:4200");
  });
});

describe("Button Test", () => {
  it('Clicks the "button", and checks for all dropdown options', () => {
    cy.visit("http://localhost:4200");

    cy.get("button").click();
    cy.contains("Login");
    cy.contains("Settings");
    cy.contains("Home");
  });
});

describe("Hyperlink Test", () => {
  it('clicking "Login" navigates to a new url', () => {
    cy.visit("http://localhost:4200");

    cy.get("button").click();
    cy.contains("Login").click();

    cy.url().should("include", "/login");
  });
});

describe("The Login Page", () => {
  it("successfully loads", () => {
    cy.visit("http://localhost:4200/login");
  });
});

describe("Box Test", () => {
  it("Tests Typing Into Boxes", () => {
    cy.visit("http://localhost:4200");

    cy.get("button").click();
    cy.contains("Login").click();

    cy.url().should("include", "/login");

    // Get an input, type into it
    cy.get("input").eq(0).type("email@com.com");
    cy.get("input").eq(0).should("have.value", "email@com.com");

    cy.get("input").eq(1).type("username");
    cy.get("input").eq(1).should("have.value", "username");

    cy.get("input").eq(2).type("password");
    cy.get("input").eq(2).should("have.value", "password");
  });
});

describe("Hyperlink Test #2", () => {
  it('clicking "Home" navigates back to home', () => {
    cy.visit("http://localhost:4200");

    cy.get("button").click();
    cy.contains("Home").click();

    cy.url().should("include", "/mainpage");
  });
});

describe("Settings Button Check", () => {
  it("Checks the settings button can be pressed", () => {
    cy.visit("http://localhost:4200");

    cy.get("button").click();
    cy.contains("Settings").click();
  });
});

describe("The Stocks Page", () => {
  it("successfully loads", () => {
    cy.visit("http://localhost:4200/stocks");
  });
});

describe("Charts Test", () => {
  it("tests charts", () => {
    cy.visit("http://localhost:4200/stocks");

    cy.get("div")
      .eq(1)
      .should("be.visible")
      .find("g.data-0 rect")
      .should("have.length", 11);
    cy.get("div").eq(3).should("be.visible");
  });
});
