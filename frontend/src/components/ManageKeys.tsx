import { useContext, useState } from "react";
import { ClientContext } from "../util/ClientContext";

const ManageKeys = () => {
  return (
    <div>
      <>Add A Key</>

      <KeyInput></KeyInput>
    </div>
  );
};

const KeyInput = () => {
  const fieldName = "Password";
  const [formState, setFormState] = useState("");
  let context = useContext(ClientContext);
  let client = context;

  const handleChange = (e: React.FormEvent<HTMLInputElement>) => {
    const input = e.currentTarget.value;
    setFormState(input);
  };

  const handleSubmit = async (e: React.SyntheticEvent) => {
    e.preventDefault();
    const response = await client.postAPIKey(formState);
    if (!response) {
      console.log("Could not update password");
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <label htmlFor="apikey-input">{`Update ${fieldName}`}</label>
        <div>
          <input
            type="apikey"
            name="apikey-input"
            id="input"
            value={formState}
            onChange={handleChange}
          />
        </div>
        <input type="submit" value="Submit" />
      </form>
    </div>
  );
};

export default ManageKeys;
