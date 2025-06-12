import { Application } from '@/interfaces/applications';
import React, { useEffect, useState } from 'react';
import api from '@/helpers/API';
import styles from './ProcessChoiceModal.module.css';
import { SaveContractForm } from '@/components/Applications/ProcessChoiceModal/SaveContract';

interface ProcessChoiceModalProps {
  selectedApp: Application;
  onClose: () => void;
  onCreateClient: () => void;
  handleContractSaved: () => void;
  onReject: () => void;
}

export const ProcessChoiceModal: React.FC<ProcessChoiceModalProps> = ({
  selectedApp,
  onClose,
  onCreateClient,
  handleContractSaved,
  onReject,
}) => {
  const [clientExists, setClientExists] = useState<boolean | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [userID, setUserID] = useState<number>(0); 
  const [showContractForm, setShowContractForm] = useState(false);

  useEffect(() => {
    const checkClient = async () => {
      try {
        setLoading(true);
        const response = await api.get(`/clients/${selectedApp.email}`);
        if (response.data.userID > 0) {
          setUserID(response.data.userID);
          setClientExists(true);
        }
      } catch (error: any) {
        if (error.response?.status === 404) {
          setClientExists(false);
        } else {
          console.error('Ошибка при проверке клиента:', error);
        }
      } finally {
        setLoading(false);
      }
    };

    if (selectedApp?.email) {
      checkClient();
    }
  }, [selectedApp]);

  return (
    <div className={styles.overlay}>
      <div className={styles.modal}>
        <h2>Обработка заявки</h2>

        <div>
          <p><strong>Организация:</strong> {selectedApp.organizationName}</p>
          <p><strong>Email:</strong> {selectedApp.email}</p>
        </div>

        {loading ? (
          <p>Проверка клиента...</p>
        ) : (
          <>
            {clientExists && showContractForm && (
              <SaveContractForm
                clientId={userID}
                applicationData={{
                  spaceCode: selectedApp.premiseNumber,
                  startDate: selectedApp.startDate,
                  endDate: selectedApp.endDate,
                }}
                onSuccess={handleContractSaved}
              />
            )}

            <div className={styles.buttonsContainer}>
              <div className={styles.leftButtons}>
                <button className={styles.danger} onClick={onReject}>
                  Отклонить
                </button>
              </div>

              <div className={styles.rightButtons}>
                {!clientExists && (
                  <button className={styles.primary} onClick={onCreateClient}>
                    Создать клиента
                  </button>
                )}
                {clientExists && !showContractForm && (
                  <button className={styles.upload} onClick={() => setShowContractForm(true)}>
                    Загрузить договор
                  </button>
                )}
                <button className={styles.secondary} onClick={onClose}>
                  Закрыть
                </button>
              </div>
            </div>
          </>
        )}
      </div>
    </div>
  );
};
