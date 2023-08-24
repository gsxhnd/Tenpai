from tensorflow.keras import Sequential
from tensorflow.keras.layers import Dense

classifier = Sequential()
classifier.add(Dense(5, activation='relu', name="test_in", input_dim=5)) # Named input
classifier.add(Dense(5, activation='relu'))
classifier.add(Dense(1, activation='sigmoid', name="test_out")) # Named output

classifier.compile(optimizer ='adam', loss='binary_crossentropy', metrics=['accuracy'])

classifier.fit([[0.1, 0.2, 0.3, 0.4, 0.5]], [[1]], batch_size=1, epochs=1);

classifier.save('examples/keras_single_input_saved_model', save_format='tf')