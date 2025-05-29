import { useEffect, useState } from "react";
import styles from './ProcessChoiceModal.module.css';
import { toast } from "react-toastify";
import api from '@/helpers/API';
import { AxiosError } from "axios";

interface RentalFormProps {
  clientId: number;
  applicationData: {
    spaceCode: number;
    startDate: string;
    endDate: string;
  };
  onSuccess?: () => void;
}

export const SaveContractForm = ({ clientId, applicationData, onSuccess }: RentalFormProps) => {
  const [file, setFile] = useState<File | null>(null);
	const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (error) {
      toast.error(error);
    }
  }, [error]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!file) return alert("Выберите файл договора.");

    const formData = new FormData();
    const payload = {
      space_code: Number(applicationData.spaceCode),
      client_id: Number(clientId),
      start_date: applicationData.startDate,
      end_date: applicationData.endDate,
    };
    formData.append("data", JSON.stringify(payload));
    formData.append("template_file", file); 
    try {
      await api.post("/rental/", formData, {});
      toast.success("Договор успешно загружен");
      onSuccess?.();
      } catch (e) {
        console.error(e);
        if (e instanceof AxiosError) {
          setError(e.response?.data.error);
        }	
      }
  };

  return (
    <form onSubmit={handleSubmit} className={styles.form}>
      <label>
        Файл договора:
        <input type="file" onChange={(e) => setFile(e.target.files?.[0] || null)} />
      </label>
      <button type="submit" className={styles.primary}>
        Загрузить договор и создать аренду
      </button>
    </form>
  );
};
